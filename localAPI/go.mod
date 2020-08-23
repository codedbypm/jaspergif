module github.com/codedbypm/jaspergif

go 1.13

replace github.com/codedbypm/jaspergif/entry => /Users/paolodev/Projects/jaspergif/entry

replace github.com/codedbypm/jaspergif/fetch => /Users/paolodev/Projects/jaspergif/fetch

replace github.com/codedbypm/jaspergif/upload => /Users/paolodev/Projects/jaspergif/upload

replace github.com/codedbypm/jaspergif/model => /Users/paolodev/Projects/jaspergif/model

replace github.com/codedbypm/jaspergif/log => /Users/paolodev/Projects/jaspergif/log

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.0.1
	github.com/codedbypm/jaspergif/entry v0.0.0-20200803220759-1b6b603cc8a8
	github.com/codedbypm/jaspergif/fetch v0.0.0-20200803220759-1b6b603cc8a8
	github.com/codedbypm/jaspergif/upload v0.0.0-20200803220759-1b6b603cc8a8
)
