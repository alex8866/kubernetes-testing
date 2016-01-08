kubernetes testings besides e2e, use go third-pary library [goconvey](https://github.com/smartystreets/goconvey).

Prerequisites:

1. This testing is focusing on kubernetes functionalities, before testing, it need to setup
an healthy kubernetes cluster. Refer to [ansible](https://github.com/kubernetes/contrib/tree/master/ansible) to do it.

2. In order to run the testing fluently, it pulls down files which will be used for testings. Run get_files.sh.
 
3. kubectl need a config file, put it in place firstly.

```
place kubectl.kubeconfig as $HOME/.kube/config

```
