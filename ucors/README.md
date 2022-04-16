# ucors?

Tool that finds CORS misconfigurations.

```sh
ucors https://target.tld/endpoint
cat urls.txt | ucors -c session=xyz123
``` 

This is a fork of [@tomnomnom's cors-blimey](https://github.com/tomnomnom/hacks/tree/master/cors-blimey).

Additional payloads:
- target.othertld (allows any tld)
- wwwxtarget.tld (failed to escape . in regex)
- target.tld{specialchar}.evil.com (host doesn't end yet)
- anything.target.tld (useful if you have xss on some subdomain)