build:
  go build -o uscis-case-status

run:
  ./uscis-case-status

init:
  go mod init uscis-case-status
  go mod tidy