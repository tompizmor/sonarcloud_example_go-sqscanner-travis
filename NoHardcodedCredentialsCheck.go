package sonarcloud_go_qscanner_travis

func noHardcodedCredentials() string  {
  password := randomElement(letters)
  return password
}
