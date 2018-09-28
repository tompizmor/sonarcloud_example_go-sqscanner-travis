package sonarcloud_go_qscanner_travis

func noHardcodedCredentials() string  {
  mypassword := randomElement(letters)
  return mypassword
}
