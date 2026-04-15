```bash
# Generate the private key for your CA
openssl genrsa -out rootCA.key 2048

# Create the Root CA certificate (Valid for ~3 years)
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 \
  -out rootCA.crt \
  -subj "/C=US/ST=State/L=City/O=AnuroopDevOps/CN=AnuroopInternalCA"
```


### Front End Certificate
```bash
# Generate private key and Certificate Signing Request (CSR)
openssl genrsa -out appgw-frontend.key 2048
openssl req -new -key appgw-frontend.key -out appgw-frontend.csr -subj "/CN=yourapp.com"

# Sign the certificate with your Root CA
openssl x509 -req -in appgw-frontend.csr -CA rootCA.crt -CAkey rootCA.key \
  -CAcreateserial -out appgw-frontend.crt -days 365 -sha256

# Export to PFX (Azure App Gateway REQUIRES PFX format for listeners)
openssl pkcs12 -export -out appgw-frontend.pfx \
  -inkey appgw-frontend.key \
  -in appgw-frontend.crt \
  -passout pass:AzurePassword123
```


### Backend Certificate
```bash
# Generate private key and CSR
openssl genrsa -out envoy-backend.key 2048
openssl req -new -key envoy-backend.key -out envoy-backend.csr -subj "/CN=envoy-ilb.internal"

# Sign the certificate with your Root CA
openssl x509 -req -in envoy-backend.csr -CA rootCA.crt -CAkey rootCA.key \
  -CAcreateserial -out envoy-backend.crt -days 365 -sha256
```