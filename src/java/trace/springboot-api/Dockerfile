FROM eclipse-temurin:21-jdk as builder

WORKDIR /usr/src/app/

COPY ./gradlew* ./settings.gradle* ./build.gradle .
COPY ./gradle ./gradle
COPY ./src ./src

RUN ./gradlew
RUN ./gradlew build

FROM eclipse-temurin:21-jre

WORKDIR /usr/src/app/

EXPOSE 8080
COPY --from=builder /usr/src/app/build/libs/*.jar ./spring-boot-application.jar
CMD [ "java", "-jar", "spring-boot-application.jar" ]
