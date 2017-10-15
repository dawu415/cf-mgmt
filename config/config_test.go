package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotalservices/cf-mgmt/config"
	. "github.com/pivotalservices/cf-mgmt/config/test_data"
	mock "github.com/pivotalservices/cf-mgmt/utils/mocks"
)

var _ = Describe("CF-Mgmt Config", func() {
	Context("Protected Org Defaults", func() {
		Describe("Defaults", func() {
			It("should setup default protected orgs", func() {
				Ω(config.DefaultProtectedOrgs).Should(HaveKey("system"))
				Ω(config.DefaultProtectedOrgs).Should(HaveKey("p-spring-cloud-services"))
				Ω(config.DefaultProtectedOrgs).Should(HaveKey("splunk-nozzle-org"))
				Ω(config.DefaultProtectedOrgs).Should(HaveLen(3))
			})
		})
	})

	Context("Default Config Reader", func() {
		Context("GetOrgConfigs", func() {
			var utilsMgrMock *mock.MockUtilsManager
			BeforeEach(func() {
				utilsMgrMock = mock.NewMockUtilsManager()
				PopulateWithTestData(utilsMgrMock)
			})
			It("should return a list of 2", func() {
				m := config.NewManager("./fixtures/config", utilsMgrMock)
				c, err := m.GetOrgConfigs()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(c).Should(HaveLen(2))
			})

			It("should return a list of 1", func() {
				m := config.NewManager("./fixtures/user_config", utilsMgrMock)
				c, err := m.GetOrgConfigs()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(c).Should(HaveLen(1))

				org := c[0]
				Ω(org.GetAuditorGroups()).Should(ConsistOf([]string{"test_org_auditors"}))
				Ω(org.GetManagerGroups()).Should(ConsistOf([]string{"test_org_managers"}))
				Ω(org.GetBillingManagerGroups()).Should(ConsistOf([]string{"test_billing_managers", "test_billing_managers_2"}))
			})

			It("should fail when given an invalid config dir", func() {
				m := config.NewManager("./fixtures/blah", utilsMgrMock)
				c, err := m.GetOrgConfigs()
				Ω(err).Should(HaveOccurred())
				Ω(c).Should(BeEmpty())
			})
		})

		Context("GetSpaceConfigs", func() {
			var utilsMgrMock *mock.MockUtilsManager
			BeforeEach(func() {
				utilsMgrMock = mock.NewMockUtilsManager()
				PopulateWithTestData(utilsMgrMock)
			})
			It("should return a single space", func() {
				m := config.NewManager("./fixtures/space-defaults", utilsMgrMock)
				cfgs, err := m.GetSpaceConfigs()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cfgs).Should(HaveLen(1))

				cfg := cfgs[0]
				Ω(cfg.Space).Should(BeEquivalentTo("space1"))
				Ω(cfg.Developer.LDAPUsers).Should(ConsistOf("default-ldap-user", "space1-ldap-user"))
				Ω(cfg.Developer.Users).Should(ConsistOf("default-user@test.com", "space-1-user@test.com"))
				Ω(cfg.Developer.LDAPGroup).Should(BeEquivalentTo("space-1-ldap-group"))

				Ω(cfg.Auditor.LDAPUsers).Should(ConsistOf("default-ldap-user", "space1-ldap-user"))
				Ω(cfg.Auditor.Users).Should(ConsistOf("default-user@test.com", "space-1-user@test.com"))
				Ω(cfg.Auditor.LDAPGroup).Should(BeEquivalentTo("space-1-ldap-group"))

				Ω(cfg.Manager.LDAPUsers).Should(ConsistOf("default-ldap-user", "space1-ldap-user"))
				Ω(cfg.Manager.Users).Should(ConsistOf("default-user@test.com", "space-1-user@test.com"))
				Ω(cfg.Manager.LDAPGroup).Should(BeEquivalentTo("space-1-ldap-group"))
			})

			It("should return a list of 2", func() {
				m := config.NewManager("./fixtures/config", utilsMgrMock)
				configs, err := m.GetSpaceConfigs()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(configs).Should(HaveLen(2))
			})

			It("should return configs for user info", func() {
				utilsMgrMock.MockFileData = map[string]interface{}{}
				PopulateWithTestData(utilsMgrMock)
				m := config.NewManager("./fixtures/user_config", utilsMgrMock)
				configs, err := m.GetSpaceConfigs()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(configs).Should(HaveLen(1))
			})

			It("should return configs for user info", func() {
				m := config.NewManager("./fixtures/user_config_multiple_groups", utilsMgrMock)
				configs, err := m.GetSpaceConfigs()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(configs).Should(HaveLen(1))
				config := configs[0]
				Ω(config.GetDeveloperGroups()).Should(ConsistOf([]string{"test_space1_developers"}))
				Ω(config.GetAuditorGroups()).Should(ConsistOf([]string{"test_space1_auditors"}))
				Ω(config.GetManagerGroups()).Should(ConsistOf([]string{"test_space1_managers", "test_space1_managers_2"}))
			})

			Context("failure cases", func() {
				It("should return an error when no security.json file is provided", func() {
					m := config.NewManager("./fixtures/no-security-json", utilsMgrMock)
					configs, err := m.GetSpaceConfigs()
					Ω(err).Should(HaveOccurred())
					Ω(configs).Should(BeNil())
				})

				It("should return an error when malformed yaml", func() {
					utilsMgrMock.MockFileDataHasError = true
					m := config.NewManager("./fixtures/bad-yml", utilsMgrMock)
					configs, err := m.GetSpaceConfigs()
					Ω(err).Should(HaveOccurred())
					Ω(configs).Should(BeNil())
				})

				It("should return an error when path does not exist", func() {
					m := config.NewManager("./fixtures/blah", utilsMgrMock)
					configs, err := m.GetSpaceConfigs()
					Ω(err).Should(HaveOccurred())
					Ω(configs).Should(BeNil())
				})
			})

		})
	})
})