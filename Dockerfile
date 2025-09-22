FROM scratch
COPY exporter /exporter
ENV MODEL_TAG=dan
ENV COUNTER_SECRET=unshackled
ENV TOKEN_SOCKET=ws://localhost:8765
EXPOSE 9100
ENTRYPOINT ["/exporter"]
