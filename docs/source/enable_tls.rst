Securing Communication With Transport Layer Security (TLS)
==========================================================

Fabric supports for secure communication between nodes using TLS.  TLS communication
can use both one-way (server only) and two-way (server and client) authentication.

Configuring TLS for peers nodes
-------------------------------

A peer node is both a TLS server and a TLS client. It is the former when another peer
node, application, or the CLI makes a connection to it and the latter when it makes
a connection to another peer node or orderer.

To enable TLS on a peer node set the following peer configuration properties:

 * ``peer.tls.enabled`` = ``true``
 * ``peer.tls.cert.file`` = fully qualified path of the file that contains the TLS server
   certificate
 * ``peer.tls.key.file`` = fully qualified path of the file that contains the TLS server
   private key

By default, TLS client authentication is turned off when TLS is enabled on a peer node.
This means that the peer node will not verify the certificate of a client (another peer
node, application, or the CLI) during a TLS handshake. To enable TLS client authentication
on a peer node, set the peer configuration property ``peer.tls.clientAuthRequired`` to
``true`` and set the ``peer.tls.clientRootCAs.files`` property to the CA chain file(s) that
contain(s) the CA certificate chain(s) that issued TLS certificates for your organization's
clients.

By default, a peer node will use the same certificate and private key pair when acting as a
TLS server and client.  To use a different certificate and private key pair for the client
side, set the ``peer.tls.clientCert.file`` and ``peer.tls.clientKey.file`` configuration
properties to the fully qualified path of the client certificate and key file,
respectively.

TLS with client authentication can also be enabled by setting the following environment
variables:

 * ``CORE_PEER_TLS_ENABLED`` = ``true``
 * ``CORE_PEER_TLS_CERT_FILE`` = fully qualified path of the server certificate
 * ``CORE_PEER_TLS_KEY_FILE`` = fully qualified path of the server private key
 * ``CORE_PEER_TLS_CLIENTAUTHREQUIRED`` = ``true``
 * ``CORE_PEER_TLS_CLIENTROOTCAS_FILES`` = fully qualified path of the CA chain file
 * ``CORE_PEER_TLS_CLIENTCERT_FILE`` = fully qualified path of the client certificate
 * ``CORE_PEER_TLS_CLIENTKEY_FILE`` = fully qualified path of the client key

When client authentication is enabled on a peer node, a client is required to send its
certificate during a TLS handshake. If the client does not send its certificate, the
handshake will fail and the peer will close the connection.

When a peer joins a channel, root CA certificate chains of the channel members are
read from the config block of the channel and are added to the TLS server and client
root CAs data structure. So, peer to peer communication, peer to orderer communication
should work seamlessly. However, if needed, you can add additional root CA certificates
by setting ``peer.tls.rootcert.file`` and ``peer.tls.clientRootCAs.files``.

Configuring TLS for orderer nodes
---------------------------------

To enable TLS on an orderer node, set the following orderer configuration properties:

 * ``General.TLS.Enabled`` = ``true``
 * ``General.TLS.PrivateKey`` = fully qualified path of the file that contains the server
   private key
 * ``General.TLS.Certificate`` = fully qualified path of the file that contains the server
   certificate

By default, TLS client authentication is turned off on orderer, as is the case with peer.
To enable TLS client authentication, set the following config property:

 * ``General.TLS.ClientAuthRequired`` = ``true``

TLS with client authentication can also be enabled by setting the following environment
variables:

 * ``ORDERER_GENERAL_TLS_ENABLED`` = ``true``
 * ``ORDERER_GENERAL_TLS_PRIVATEKEY`` = fully qualified path of the file that contains the
   server private key
 * ``ORDERER_GENERAL_TLS_CERTIFICATE`` = fully qualified path of the file that contains the
   server certificate
 * ``ORDERER_GENERAL_TLS_CLIENTAUTHREQUIRED`` = ``true``

   When an orderer participates in a channel, root CA certificate chains of the channel members are
   read from the config block of the channel and are added to the TLS server and client
   root CAs data structure. So, orderer to orderer communication
   should work seamlessly. However, if needed, you can add additional root CA certificates
   by setting ``General.TLS.RootCAs`` and ``General.TLS.ClientRootCAs``.

Configuring TLS for the peer CLI
--------------------------------

The following environment variables must be set when running peer CLI commands against a
TLS enabled peer node:

* ``CORE_PEER_TLS_ENABLED`` = ``true``
* ``CORE_PEER_TLS_ROOTCERT_FILE`` = fully qualified path of the file that contains cert chain
  of the CA that issued the TLS server cert

If TLS client authentication is also enabled on the remote server, the following variables
must to be set in addition to those above:

* ``CORE_PEER_TLS_CLIENTAUTHREQUIRED`` = ``true``
* ``CORE_PEER_TLS_CLIENTCERT_FILE`` = fully qualified path of the client certificate
* ``CORE_PEER_TLS_CLIENTKEY_FILE`` = fully qualified path of the client private key

When running a command that connects to orderer service, like `peer channel <create|update|fetch>`
or `peer chaincode <invoke>`, following command line arguments must also be specified
if TLS is enabled on the orderer:

* --tls
* --cafile <fully qualified path of the file that contains cert chain of the orderer CA>

If TLS client authentication is enabled on the orderer, the following arguments must be specified
as well:

* --clientauth
* --keyfile <fully qualified path of the file that contains the client private key>
* --certfile <fully qualified path of the file that contains the client certificate>

TLS with proxy servers
----------------------

Since Fabric components verify each other with TLS, if you are using a proxy server,
it must be configured for TLS passthrough (non-terminating) so that the TLS
credentials are passed on to the Fabric components.

Subject Alternative Names
-------------------------

Each TLS server must have one or more Subject Alternative Names configured in its TLS certificate that specifies its domain name or IP address.
When a TLS client attempts to connect to the TLS server, it verifies that one of the Subject Alternative Names matches the address it is trying to connect to.

When creating a TLS certificate, the Subject Alternative Name(s) must be specified.
If using Fabric CA to create a TLS certificate, specify the Subject Alternative Names as a comma-separated list in the ``--csr.hosts`` flag in the enroll command.
If using cryptogen to create a TLS certificate, specify the Subject Alternative Names as a list under the ``SANS`` element of the cryptogen config yaml.

Debugging TLS issues
--------------------

If you see the error message ``remote error: tls: bad certificate`` on the server side
(for example on the peer node or ordering service node when making requests from a client),
it usually means that the client is not configured to trust the signer of the server's TLS certificate.
Check the client's ``CORE_PEER_TLS_ROOTCERT_FILE`` (for connections to peer nodes)
or ``--cafile`` (for connections to orderer nodes).
The corresponding error on the client side in these cases is the handshake error ``x509: certificate signed by unknown authority``
and ultimately connection failure with ``context deadline exceeded``.
The problem may also be a Subject Alternative Names issue. In these cases the the handshake error on the client side will be
``tls: failed to verify certificate: x509: certificate is valid for <configured_SAN>, not <attempted_address>``.

If you see the error message ``remote error: tls: bad certificate`` on the client side, it
usually means that the TLS server has enabled client authentication and the server either did
not receive the correct client certificate or it received a client certificate that it does
not trust. Make sure the client is sending its certificate and that it has been signed by one
of the CA certificates trusted by the peer or orderer node.

To receive additional debug information, enable ``GRPC debug`` on both the TLS client
and the server side to get additional information. To enable ``GRPC debug``, set the
environment variable ``FABRIC_LOGGING_SPEC`` to include ``grpc=debug``. For example, to
set the default logging level to ``INFO`` and the GRPC logging level to ``DEBUG``, set
the logging specification to ``grpc=debug:info``.

You can check a TLS certificate against a trusted CA certificate by using the "openssl verify" command.

.. Licensed under Creative Commons Attribution 4.0 International License
   https://creativecommons.org/licenses/by/4.0/
