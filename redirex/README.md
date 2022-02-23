# redirex

This tool generates a wordlist for open redirects.

```sh
go install github.com/wfinn/stuff/redirex@latest
redirex -target target.tld -attacker attacker.tld
redirex -h # to see more options
```
---

To find some more redirects you can use these parameters at login/logout.

```
?Redirect=/Redirect&RedirectUrl=/RedirectUrl&ReturnUrl=/ReturnUrl&Url=/Url&action=/action&action_url=/action_url&backurl=/backurl&burl=/burl&callback_url=/callback_url&checkout_url=/checkout_url&clickurl=/clickurl&continue=/continue&data=/data&dest=/dest&destination=/destination&desturl=/desturl&ext=/ext&forward=/forward&forward_url=/forward_url&go=/go&goto=/goto&image_url=/image_url&jump=/jump&jump_url=/jump_url&link=/link&linkAddress=/linkAddress&location=/location&login=/login&logout=/logout&next=/next&origin=/origin&originUrl=/originUrl&page=/page&pic=/pic&q=/q&qurl=/qurl&recurl=/recurl&redir=/redir&redirect=/redirect&redirect_uri=/redirect_uri&redirect_url=/redirect_url&request=/request&return=/return&returnTo=/returnTo&return_path=/return_path&return_to=/return_to&rit_url=/rit_url&rurl=/rurl&service=/service&sp_url=/sp_url&src=/src&success=/success&target=/target&u=/u&u1=/u1&uri=/uri&url=/url&view=/view
```

To test for for redirects that require full urls use `echo "params" | sed 's#=/#=https://TARGETDOMAIN/#g'`.

---

Uses:
- https://github.com/dbzer0/ipfmt MIT License
- https://github.com/jpillora/go-tld MIT License
