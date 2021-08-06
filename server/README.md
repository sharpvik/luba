# Luba Server

Luba Server is a [revealing](#revealing) load balancer. It sits there, waiting
for your services to connect to it (see [API.md](API.md) `POST /hey`), and as
soon as they do, it will be able to share their addresses with the world in a
way that evenly spreads traffic between the instances.

<a name="revealing"></a> _Revealing_ load balancers don't proxy your traffic
but rather serve as a discovery service for anyone who wants to connect.

## Why Revealing Load Balancer?

In most cases, it is ok to have a central load balancer that proxies traffic to
multiple instances. However, there is a use case where this is impossible or at
least impractical: high-troughput microservices.

Sometimes, the reason you run multiple instances of a service is because you
expect each one of them to handle _a lot of traffic_ and also, _fast_. In that
case, having a single load balancer node in the middle is not feasible, because
if you can afford to run such a high-throughput server to proxy each one of
those miriads of requests, maybe just run your service on there instead.

For those of us who don't have that kind of technological superiority, we'd
want a load balancer that doesn't serve as a redirection node but simply gives
out the final address of the service. Then, the client can just communicate
with that instance directly.
