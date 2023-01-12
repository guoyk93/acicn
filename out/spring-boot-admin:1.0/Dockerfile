FROM dev-build/acicn/jdk:8-debian-11
WORKDIR /work
RUN curl -sSL -o spring-boot-admin.jar "https://github.com/guoyk93/acicn-assets/releases/download/v20221122/spring-boot-admin-1.0.0-RELEASE.jar"
CMD ["java-wrapper", "-jar", "spring-boot-admin.jar"]