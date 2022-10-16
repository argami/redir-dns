# Redir DNS

Is a [Caddy](https://caddyserver.com/) Module to create redirections based in a TXT record for a domain.

Using caddy allows to make https redirections activating the auto https.

### How to use

To include in Caddy needs to recompile caddy including the module

```bash
$ xcaddy build --with github.com/argami/redir-dns@0174c1a
```

##### - Usage:
`redir_dns [matcher]`

***Caddyfile Example:***

```caddyfile
{
  order redir_dns after redir # Required
}

localhost {
  redir_dns # if matcher
}
```

##### - DNS Configuration

Yo need to create a TXT record for the domain with `_redirdns` value for example to redirect www.myhost.com to anotherhost.com your DNS config should be like:

```
CNAME www.arthomecarcasonne.fr. 1m00s   "mycaddyhost.com."
TXT _redirdns.www.myhost.com. 1m00s   "https://anotherhost.com"
```

