module github.com/codedbypm/jaspergify

go 1.13

replace github.com/codedbypm/jaspergify/entry => /Users/paolodev/Projects/jaspergify/entry

replace github.com/codedbypm/jaspergify/fetch => /Users/paolodev/Projects/jaspergify/fetch

replace github.com/codedbypm/jaspergify/upload => /Users/paolodev/Projects/jaspergify/upload

replace github.com/codedbypm/jaspergify/model => /Users/paolodev/Projects/jaspergify/model

replace github.com/codedbypm/jaspergify/log => /Users/paolodev/Projects/jaspergify/log

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.0.1
	github.com/codedbypm/jaspergify/entry v0.0.0-20200803220759-1b6b603cc8a8
	github.com/codedbypm/jaspergify/fetch v0.0.0-20200803220759-1b6b603cc8a8
	github.com/codedbypm/jaspergify/upload v0.0.0-20200803220759-1b6b603cc8a8
)
