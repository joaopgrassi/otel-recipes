FROM otel/opentelemetry-collector:0.40.0
COPY conf.yml .
EXPOSE 1888
EXPOSE 8888
EXPOSE 8889
EXPOSE 13133
EXPOSE 4317
EXPOSE 55670
CMD [ "--config=conf.yml" ]