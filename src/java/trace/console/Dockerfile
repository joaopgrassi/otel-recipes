FROM gradle:7.3.3-jdk17-alpine
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle build
CMD [ "java", "-jar", "app/build/libs/app.jar" ]
