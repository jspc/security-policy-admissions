# Security Policy Admissions

A simple validating admissions controller/ webhook to deny deployment of anything which:

1. Has `priviledged: true`
2. Has `allowEscalation: true`
3. Has `runAsNonRoot: false`
4. Has `readOnlyRootFilesystem: false`
5. Doesn't drop all capabilities

This will run against any deployment, including against third party deployments via helm, to gate my Kubernetes clusters in the way I want them

## Why not `AdmissionConfiguration` (or some variation thereof)

Mainly I'd like to be a lot stricter, and a lot more flexible, than the out-of-the-box stuff allows. I also want to be able to run against things other than pods in the future, such as my own CRDs

## Development

This project uses pretty standard... everything, really. The usual `go build` and `go test` works just fine.

## Licence

Copyright 2023 James Condron

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
