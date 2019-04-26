FROM grafana/grafana:6.0.2
MAINTAINER Oliver Nadj <mr.oliver.nadj@gmail.com>

ADD --chown=grafana:grafana grafana.db /var/lib/grafana/grafana.db

ENTRYPOINT [ "/run.sh" ]