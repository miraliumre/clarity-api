# Clarity API

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