package handlers_test

import (
	"bosh-dns/dns/server/handlers"
	"bosh-dns/dns/server/handlers/handlersfakes"
	"bosh-dns/dns/server/internal/internalfakes"
	"github.com/miekg/dns"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ bool = Describe("CacheHandler", func() {
	var (
		cacheHandler   handlers.CachingDNSHandler
		fakeWriter     *internalfakes.FakeResponseWriter
		fakeDnsHandler *handlersfakes.FakeDNSHandler
	)

	BeforeEach(func() {
		fakeDnsHandler = &handlersfakes.FakeDNSHandler{}
		fakeWriter = &internalfakes.FakeResponseWriter{}
		cacheHandler = handlers.NewCachingDNSHandler(fakeDnsHandler)
	})

	Describe("ServeDNS", func() {
		Context("when the request doesn't have recursion desired bit set", func() {
			It("responds with rcode refused", func() {
				m := &dns.Msg{}
				m.SetQuestion("my-instance.my-group.my-network.my-deployment.bosh.", dns.TypeANY)
				m.RecursionDesired = false
				cacheHandler.ServeDNS(fakeWriter, m)
				message := fakeWriter.WriteMsgArgsForCall(0)
				Expect(message.Rcode).To(Equal(dns.RcodeRefused))
			})
		})
	})
})
