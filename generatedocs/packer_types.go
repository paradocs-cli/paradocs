package generate_docs

type PackerAzureSourceBlock struct {
	ImageData struct {
		Offer     string
		Sku       string
		Publisher string
	}
	AzureAuth struct {
		SubscriptionId string
		ClientId       string
		ClientSecret   string
		TenantId       string
	}
	SshConfig struct {
		SshUsername   string
		SshPassword   string
		SshPublicKey  string
		SshPrivateKey string
	}
	ManagedImageConfig struct {
		ManagedImageName  string
		ManageImageRgName string
	}
	Tags         map[string]string
	Communicator string
	Location     string
	OsType       string
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

type BuildBlock struct{}
