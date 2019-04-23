FROM scratch
WORKDIR /profilesvcapp
COPY profileservice .
CMD profileservice