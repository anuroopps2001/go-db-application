import requests
import sys
import os
from azure.identity import DefaultAzureCredential


TARGET_NAMESPACE = os.getenv("TARGET_NAMESPACE", "go-db-app-dev")
BASE_URL = os.getenv("PROMETHEUS_URL", "https://go-app-prometheus-geacauf4ewaxgrd7.southindia.prometheus.monitor.azure.com").rstrip('/')
PROMETHEUS_URL = f"{BASE_URL}/api/v1/query"
CANARY_VERSION= "02"
ERROR_THRESHOLD=2.0
LATENCY_THRESHOLD=0.2

def get_azure_token():
    """Fetches a Bearer token for Azure Managed Prometheus."""
    try:
        # DefaultAzureCredential works in local Dev (CLI login) and AKS (Managed Identity)
        credential = DefaultAzureCredential()
        # The scope for Managed Prometheus is fixed
        token = credential.get_token("https://prometheus.monitor.azure.com/.default")
        return token.token
    except Exception as e:
        print(f"Authentication Error: {e}")
        sys.exit(1)

def query_prometheus(query):
    token = get_azure_token()
    headers = {"Authorization": f"Bearer {token}"}

    try:
        response = requests.get(PROMETHEUS_URL, params={'query': query}, headers=headers, timeout=5)
        response.raise_for_status() # Catch 401/403 errors
        data = response.json()
        results = data.get('data', {}).get('result', []) # List of dictionaries
        
        if not results:
            return None
        
        return float(results[0]['value'][1])
    
    except Exception as e:
        print(f"Error querying prometheus: {e}")
        return None
    

def evaluate_query():
    print(f"\n--- Evaluating Canary Version {CANARY_VERSION} ---")

    # 1. Query Error Rate %
    error_query = f'''
    (
      sum(rate(http_requests_total{{namespace="{TARGET_NAMESPACE}", status=~"4..|5..", version="{CANARY_VERSION}"}}[5m]))
      / 
      (sum(rate(http_requests_total{{namespace="{TARGET_NAMESPACE}", version="{CANARY_VERSION}"}}[5m])) or vector(1))
    ) * 100 or vector(0)'''
    error_rate = query_prometheus(error_query)


    # 2. Query P95 Latency (Seconds)
    latency_query = f'''
    histogram_quantile(0.95,
        sum by (le) (
            rate(http_request_duration_seconds_bucket{{namespace="{TARGET_NAMESPACE}", version="{CANARY_VERSION}"}}[5m])
        )
    ) or vector(0)'''
    p95_latency = query_prometheus(latency_query)

    # --- EVALUATION ---
    if error_rate is None or p95_latency is None:
        print(f"Failed to fetch metrics")
        sys.exit(1)
    
    print(f"Current Environment: {TARGET_NAMESPACE}")
    print(f"Error Rate: {error_rate:.2f}% (Threshold: {ERROR_THRESHOLD}%)")
    print(f"P95 Latency: {p95_latency:.3f} (Threshold: {LATENCY_THRESHOLD}s)")

    error_pass = error_rate <=ERROR_THRESHOLD
    latency_pass = p95_latency <= LATENCY_THRESHOLD


    if error_pass and latency_pass:
        print("\nResult: PASS")
        sys.exit(0)

    else:
        print("\nResult: FAIL")
        if not error_pass: 
            print(f"- Reason: Error rate ({error_rate:.2f}%) exceeds threshold")
        if not latency_pass: 
            print(f"- Reason: Latency ({p95_latency:.3f}s) exceeds threshold")
        sys.exit(1)

if __name__ == "__main__":
    evaluate_query()