# Clarity API

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=miraliumre_clarity-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=miraliumre_clarity-api) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=miraliumre_clarity-api&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=miraliumre_clarity-api) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=miraliumre_clarity-api&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=miraliumre_clarity-api)

The Clarity API is a subsystem of Clarity, a solution developed by Miralium
Research for document validation.

Primarily designed to address the compliance needs of users in cybersecurity
and auditing domains, Clarity ensures that the documents shared amongst the
involved parties are untouched and exactly as created by Miralium Research.

## Endpoints

- `/v1/documents/:hash`: Validate a document via hash value (e.g., SHA-256)
- `/v1/certs`: Provides an array of valid X509 certificate chains for
  validating digitally signed documents

## License

This project is distributed under [The Unlicense](LICENSE.txt).