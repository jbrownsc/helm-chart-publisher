## How to use

0. Copy values.yaml to local-values.yaml
0. Override any local settings or set via command line
  0. NOTES.txt will guide you on how to access via service or ingress
0. Install chart
  `helm install . --namespace XXX --name helm-chart-publisher -f local-values.yaml --set aws_access_key=XXX,aws_secret_key=XXX`
0. Package this chart
  `helm package .`
0. Test posting a chart
  `curl -i -X PUT -F repo=internal -F chart=@$(ls -1tr *.tgz | tail -n1) http://<service or ingress>/charts`
0. Browse to http://<service or ingress>/internal/index.yaml
