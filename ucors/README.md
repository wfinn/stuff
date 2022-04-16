# ucors?

Tool that finds CORS misconfigurations.

```sh
echo https://target.tld/endpoint | ucors
ucors < urls.txt
``` 

This is a fork of [@tomnomnom's cors-blimey](https://github.com/tomnomnom/hacks/tree/master/cors-blimey).

Additional features:
- target.othertld
- target.tld{specialchar}.evil.com
- anything.target.tld (useful if you have xss on some subdomain)