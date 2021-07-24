## OpenSSL
### Generate private RSA key
To generate RSA using OpenSSL tool, we need to use `genrsa` command to below,

```bash
openssl genrsa -out server.key 2048
```
1. Specifies which algorithm to use to create the key. OpenSSL supports creating keys with a different algorithm like RSA, DSA, ECDSA. ALL types are practical for use in all scenarios. For example, for web server keys commonly uses RSA. in our case, we need to generate RSA type key.
2. Specifies the name of the generated key. Can have any name with `.key` as extension.
3. Specifies the size of the key. The default size for RSA keys is only 512 bits, which is not secure because an intruder can use brute force to recover you private key. So we use a 2048-bit RSA key which is considered to be secure.


### Generate public key/certificate
```bash
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields, there will be a default value 
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:US
State or Province Name (full name) []:California
Locality Name (eg, city) []:Mountain View
Organization Name (eg, company) []:O’Reilly Media, Inc
Organizational Unit Name (eg, section) []:Publishers
Common Name (eg, fully qualified hostname[]:localhost
Email Address []:webmaster@localhost

```
1. Specifies the format of the public certificate X.509 is a standard format which is used in many Internet protocols, including TLS/SSL.
2. Specifies the secure hash algorithm.
3. Specifies private key(`server.key`)file location which we generated before.
4. Specifies the name of generated certificate. Can have any name with `.crt` as extension.
5. Specifies the lifetime of the certificate to 3650 dats.

```bash
# Quickly generate private key and certificate with a single command.
openssl req -x509 -sha256 -nodes -days 3650 -newkey rsa:4096 -keyout server.key -out server.crt -subj "/C=US/ST=California/L=Mountain View/O=O’Reilly Media, Inc/OU=Publishers/CN=localhost/emailAddress=webmaster@localhost"

# View certificate infomation
openssl x509 -text -in server.crt -noout
```
### [NOTE]
The most important question among the questions asked when generating the certificate, is the “Common Name” which is composed of the host, domain, or IP address of the server related to the certificate. This name is used during the verification and if the host name doesn’t match the common name a warning is raised.

In order to secure java application, we need to provide key store (`.pem file`). We can easily convert the server and client keys using following command.
```bash
$ openssl pkcs8 -topk8 -inform pem -in server.key -outform pem -nocrypt -out server.pem
$ openssl pkcs8 -topk8 -inform pem -in client.key -outform pem -nocrypt -out client.pem
```

## mkcert

## certstrap