package generate_docs

type PackerAzureSourceBlock struct {
	ImageData struct {
		Offer     string
		Sku       string
		Publisher string
	}
	Tags              map[string]string
	SubscriptionId    string
	ClientId          string
	ClientSecret      string
	TenantId          string
	Communicator      string
	ManagedImageName  string
	ManageImageRgName string
	Location          string
	OsType            string
	SshUsername       string
	SshPassword       string
	SshPublicKey      string
	SshPrivateKey     string
}

type PackerAwsSourceBlock struct {
	ImageData struct {
		Offer     string
		Sku       string
		Publisher string
	}
	Tags              map[string]string
	SubscriptionId    string
	ClientId          string
	ClientSecret      string
	TenantId          string
	Communicator      string
	ManagedImageName  string
	ManageImageRgName string
	Location          string
	OsType            string
	SshUsername       string
	SshPassword       string
	SshPublicKey      string
	SshPrivateKey     string
}

type PackerGcpSourceBlock struct {
	ImageData struct {
		Offer     string
		Sku       string
		Publisher string
	}
	Tags              map[string]string
	SubscriptionId    string
	ClientId          string
	ClientSecret      string
	TenantId          string
	Communicator      string
	ManagedImageName  string
	ManageImageRgName string
	Location          string
	OsType            string
	SshUsername       string
	SshPassword       string
	SshPublicKey      string
	SshPrivateKey     string
}
