FROM eclipse-temurin:21-jdk as builder

WORKDIR /usr/src/app/

COPY ./gradlew* ./settings.gradle* ./app/build.gradle .
COPY ./gradle ./gradle

RUN ./gradlew
RUN ./gradlew build

COPY ./app/build/libs/app.jar ./

FROM eclipse-temurin:21-jre

WORKDIR /usr/src/app/

COPY --from=builder /usr/src/app/ ./
CMD [ "java", "-jar", "app.jar" ]
