# dynamic-servers-nginx

A ideia do repositório é termos uma validação da seguinte situaçao.
Em um nginx que serve como api gateway simples (proxy reverse, redirect, cache, etc.) como ele iria trablhar com uma api que escala por exemplo, ou um loadbalancer da AWS que pode mudar os ips que são resolvidos no dns dela.

*A implementação da solução é baseada neste artigo: [https://distinctplace.com/2017/04/19/nginx-resolver-explained/](https://distinctplace.com/2017/04/19/nginx-resolver-explained/)
*

## Soluções

Conforme mencionado no artigo, nós temos 4 soluções:

1. Um processo que vigia o dns e sempre que ele mudar rodar um `nginx -s reload`.
2. Utilizar o Nginx Plus que prove a directiva `resolve` que faz essa modificação do ip na memoria compartilhada do nginx. O problema é que o Nginx Plus tem um custo de ~U$5K por instancia anualmente.
3. Utilizar o modulo customizado [nginx-upstream-dynamic-servers](https://github.com/GUI/nginx-upstream-dynamic-servers). Porém ele tem muito tempo que não recebe atualizações (O ultimo commit foi à 6 anos).
4. Utilizar uma variavel e setar a directiva resolver com o dns server do ambiente. No exemplo deste repositório foi utilizado o dns server do docker.

## Implementação

```nginx
# O pulo do gato está aqui, eu peguei este ip subindo um container e fazendo um cat /etc/resolve.conf
# Porém se vc estiver na aws, vc pode setar os ips 10.0.0.2, 169.254.169.253 que são os dns servers padrão.
# https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-limits
resolver 127.0.0.11 valid=10s;

server {
    listen       80;
    server_name  localhost;

    set $upstream_endpoint api:8080;

    location / {
        proxy_pass http://$upstream_endpoint;
        proxy_redirect off;
    }
}
```

## Teste

```bash
docker-compose up --build
docker-compose scale api=3  
```
