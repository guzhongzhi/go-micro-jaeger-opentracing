FROM alpine
ADD test4-srv /test4-srv
ENTRYPOINT [ "/test4-srv" ]
