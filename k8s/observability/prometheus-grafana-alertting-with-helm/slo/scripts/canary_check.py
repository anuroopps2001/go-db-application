import requests
import sys

PROMETHEUS_URL= "" # PrometheusURL
CANARY_VERSION= "02"
ERROR_THRESHOLD="2.0"
LATENCY_THRESHOLD="0.5"

def query_prometheus(query):
    try:
        response = requests.get(PROMETHEUS_URL, params={'query': query})
        data = response.json()
        results = data.get('data', {}).get('result', []) # List of dictionaries
        if not results:
            return 0.0
        return float(results[0]['value'][1])
    
    except Exception as e:
        print(f"Error querying prometheus: {e}")
        return None
    

def evaluate_query():
    print(f"\n--- Evaluating Canary Version {CANARY_VERSION} ---")

    # 1. Query Error Rate %
    error_query = f'''
    (
     sum(rate(http_requests_total{{status=~"4..|5..",version="{CANARY_VERSION}"}}[5m]))
     / 
     sum(rate(http_requests_total{{version="{CANARY_VERSION}"}}[5m])
    ) * 100) or vector(0)'''
    error_rate = query_prometheus(error_query)


    # 2. Query P95 Latency (Seconds)
    latency_query = f'''
    histogram_quantile(0.95,
       sum by (le) (
       rate(http_request_duration_seconds_bucket{{version="{CANARY_VERSION}"}}[5m])
    ))'''
    p95_latency = query_prometheus(latency_query)

    # --- EVALUATION ---
    if error_rate is None or p95_latency is None:
        print(f"Failed to fetch metrics")
        sys.exit(1)
    
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
            print("- Reason: Error rate too high")
        if not latency_pass: 
            print("- Reason: Latency exceeds threshold")
        sys.exit(1)

if __name__ == "__main__":
    evaluate_query()