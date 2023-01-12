FROM {{Lookup .upstream}}

WORKDIR /work

RUN curl -sSL -o spring-boot-admin.jar "{{.spring_boot_admin_download_url}}"

CMD ["java-wrapper", "-jar", "spring-boot-admin.jar"]