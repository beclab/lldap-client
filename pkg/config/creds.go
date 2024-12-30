package config

import (
	corev1 "k8s.io/api/core/v1"
)

type Secret struct {
	*corev1.Secret
}

func (s *Secret) GetUsername(key string) string {
	return string(s.Data[key])
}

func (s *Secret) GetPassword(key string) string {
	return string(s.Data[key])
}

func (s *Secret) GetLLdapHost(key string) string {
	return string(s.Data[key])
}

func NewConfigFromSecret(secret *Secret) *Config {
	return &Config{
		Host:     secret.GetLLdapHost("lldap-host"),
		Username: secret.GetUsername("lldap-ldap-user-dn"),
		Password: secret.GetPassword("lldap-ldap-user-pass"),
	}
}
