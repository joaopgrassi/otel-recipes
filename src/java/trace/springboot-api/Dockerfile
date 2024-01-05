FROM gradle:jdk21-alpine as build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle build --no-daemon
CMD [ "java", "-jar", "app/build/libs/app.jar" ]


FROM openjdk:21-slim
EXPOSE 8080
RUN mkdir /app
COPY --from=build /home/gradle/src/build/libs/*.jar /app/spring-boot-application.jar
ENTRYPOINT ["java", "-jar","/app/spring-boot-application.jar"]
