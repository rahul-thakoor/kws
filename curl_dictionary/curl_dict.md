# Using cURL To Query An Online Dictionary

## Introduction

We often face situations where we need to verify the meaning of a word or simply translate it. Being able to verify a word in the dictionary without opening a browser can be very handy.  In this article, we are going to demonstrate how to use the command line tool `curl` to query an online dictionary such as `dict.org`.

### cURL
cURL[1] is a software project providing a library,`libcurl` and the command line tool `curl`. `curl` is used to transfer data to or from a server and supports numerous protocols such as DICT, FILE, FTP, FTPS, GOPHER, HTTP, HTTPS,  IMAP,
IMAPS,  LDAP,  LDAPS,  POP3,  POP3S,  RTMP, RTSP, SCP, SFTP, SMB, SMBS,
SMTP, SMTPS, TELNET and TFTP. 

### DICT Protocol
DICT is a dictionary network protocol created by the DICT Development Group.[2] The protocol operates on TCP port 2628.

### dict.org
`dict.org` is a website that hosts several dictionaries and understands the DICT dictionary protocol. It also provides a web interface to query the dictionaries.

## Query a word
Look up the definition of a word by using the command:

```shell
    curl dict://dict.org/d:<word>
```

`curl` can automatically communicate with `dict.org` using the DICT protocol, so the `dict://` can be omitted in queries.

e.g Defining `telephone`:

```shell
curl dict.org/d:telephone

220 pan.alephnull.com dictd 1.12.1/rf on Linux 4.4.0-1-amd64 <auth.mime> <12788333.7639.1519406212@pan.alephnull.com>
250 ok
150 2 definitions retrieved
151 \"Telephone\" gcide \"The Collaborative International Dictionary of English v.0.48\"
Telephone \Tel\"e*phone\, n. [Gr. ? far off + ? sound.] (Physics)
   An instrument for reproducing sounds, especially articulate
   speech, at a distance.
   [1913 Webster]

   Note: The ordinary telephone consists essentially of a device
         by which currents of electricity, produced by sounds
         through the agency of certain mechanical devices and
         exactly corresponding in duration and intensity to the
         vibrations of the air which attend them, are
         transmitted to a distant station, and there, acting on
         suitable mechanism, reproduce similar sounds by
         repeating the vibrations. The necessary variations in
         the electrical currents are usually produced by means
         of a microphone attached to a thin diaphragm upon which
         the voice acts, and are intensified by means of an
         induction coil. In the magnetic telephone, or
         magneto-telephone, the diaphragm is of soft iron placed
         close to the pole of a magnet upon which is wound a
         coil of fine wire, and its vibrations produce
         corresponding vibrable currents in the wire by
         induction. The mechanical, or string, telephone is a
         device in which the voice or sound causes vibrations in
         a thin diaphragm, which are directly transmitted along
         a wire or string connecting it to a similar diaphragm
         at the remote station, thus reproducing the sound. It
         does not employ electricity.
         [1913 Webster]
.
151 \"Telephone\" gcide \"The Collaborative International Dictionary of English v.0.48\"
Telephone \Tel\"e*phone\, v. t.
   To convey or announce by telephone.
   [1913 Webster]
.
250 ok [d/m/c = 2/0/18; 0.000r 0.000u 0.000s]
221 bye [d/m/c = 0/0/0; 0.000r 0.000u 0.000s]

```

There can be cases where no definitions are found:

```shell
220 pan.alephnull.com dictd 1.12.1/rf on Linux 4.4.0-1-amd64 <auth.mime> <12788573.11258.1519406381@pan.alephnull.com>
250 ok
552 no match [d/m/c = 0/0/148; 0.000r 0.000u 0.000s]
221 bye [d/m/c = 0/0/0; 0.000r 0.000u 0.000s]
```

## Refining searches
Querying `dict.org` directly returns results from the first dictionary in which the query matches. In the above example, the dicitonary is `The Collaborative International Dictionary of English v.0.48`.



## References

[1] https://en.wikipedia.org/wiki/CURL

[2] https://en.wikipedia.org/wiki/DICT
