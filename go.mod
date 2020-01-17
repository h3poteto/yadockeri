module github.com/h3poteto/yadockeri

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3 // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.20.0+incompatible // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/chai2010/gettext-go v0.0.0-20170215093142-bf70f2a70fb1 // indirect
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.1-0.20190620180102-5e25c22bd5d6+incompatible // indirect
	github.com/docker/distribution v0.0.0-20190711223531-1fb7fffdb266 // indirect
	github.com/docker/docker v1.4.2-0.20190327010347-be7ac8be2ae0 // indirect
	github.com/emicklei/go-restful v1.1.4-0.20170410110728-ff4f55a20633 // indirect
	github.com/evanphx/json-patch v4.2.0+incompatible // indirect
	github.com/fatih/color v1.7.1-0.20181010231311-3f9d52f7176a // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/go-github/v27 v27.0.6
	github.com/google/go-querystring v1.0.1-0.20190318165438-c8c88dbee036 // indirect
	github.com/gorilla/context v1.1.2-0.20190627024605-8559d4a6b87e // indirect
	github.com/gorilla/securecookie v1.1.2-0.20190707033817-86450627d8e6 // indirect
	github.com/gorilla/sessions v1.2.0
	github.com/gosuri/uitable v0.0.4
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/imdario/mergo v0.3.8-0.20190722133502-4c317f2286be // indirect
	github.com/jessevdk/go-assets v0.0.0-20160921144138-4f4301a06e15
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo-contrib v0.0.0-20180222075343-7d9d9632a4aa
	github.com/labstack/gommon v0.2.10-0.20190711173550-1f43c1e9c3eb // indirect
	github.com/lib/pq v1.2.0
	github.com/mattn/go-colorable v0.1.3-0.20190708054220-c52ace132bf4 // indirect
	github.com/mattn/go-isatty v0.0.9-0.20190708054338-da60ac76bf70 // indirect
	github.com/mattn/go-runewidth v0.0.5-0.20181218000649-703b5e6b11ae // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1.0.20190228220655-ac19fd6e7483 // indirect
	github.com/pkg/errors v0.8.2-0.20190227000051-27936f6d90f9
	github.com/rubenv/sql-migrate v0.0.0-20190717103323-87ce952f7079 // indirect
	github.com/russross/blackfriday v1.5.3-0.20190616195246-a925a152c144 // indirect
	github.com/sergi/go-diff v1.0.1-0.20180205163309-da645544ed44 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/valyala/bytebufferpool v1.0.1-0.20180905182247-cdfbe9377474 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/net v0.0.0-20191112182307-2180aed22343
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/grpc v1.18.0 // indirect
	gopkg.in/gorp.v1 v1.7.2 // indirect
	gopkg.in/square/go-jose.v2 v2.1.7-0.20180411045311-89060dee6a84 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.2.7
	k8s.io/client-go v0.17.1
	k8s.io/helm v2.10.0-rc.2.0.20190724233313-8cacb2c35df1+incompatible
	k8s.io/kubernetes v1.15.0 // indirect
	vbom.ml/util v0.0.0-20180919145318-efcd4e0f9787 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190620084959-7cf5895f2711
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190620085554-14e95df34f1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190620085212-47dc9a115b18
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190620085706-2090e6d8f84c
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191016115248-b061d4666016
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115051-4323e76404b0
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190620085130-185d68e6e6ea
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190817025403-3ae76f584e79
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191016115443-72c16c0ea390
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112329-27bff66d0b7c
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114902-c7514f1b89da
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114328-7650d5e6588e
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114710-682e84547325
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114520-100045381629
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115707-22244e5b01eb
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191016113728-f445c7b35c1c
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112728-ceb381866e80
)
