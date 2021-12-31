# PlantUML C4 to object parser Contribution Guide

The owners of this repository welcomes your contribution. To make the process as seamless as possible, we recommend you read this contribution guide.

## Development Workflow

Start by forking the this GitHub repository, make changes in a branch and then send a pull request. We encourage pull requests to discuss code changes. Here are the steps in details:

### Setup your puml-c4-to-object-go GitHub Repository
Fork [puml-c4-to-object-go](https://github.com/rzrbld/puml-c4-to-object-go/fork) upstream source repository to your own personal repository.

### Set up git remote as ``upstream``
```sh
$ git clone https://github.com/rzrbld/puml-c4-to-object-go
$ cd puml-c4-to-object-go
$ git remote add upstream https://github.com/rzrbld/puml-c4-to-object-go
$ git fetch upstream
$ git merge upstream/main
...
```

### Create your feature branch
Before making code changes, make sure you create a separate branch for these changes

```
$ git checkout -b my-new-feature
```

### Test changes
After your code changes, make sure
- To run `go build c4d.go` and see no errors
- To squash your commits into a single commit. `git rebase -i`. It's okay to force update your pull request.
- To run some tests.

### Commit changes
After verification, commit your changes. This is a [great post](https://chris.beams.io/posts/git-commit/) on how to write useful commit messages

```
$ git commit -am 'Add some feature'
```

### Push to the branch
Push your locally committed changes to the remote origin (your fork)
```
$ git push origin my-new-feature
```

### Create a Pull Request
Pull requests can be created via GitHub. Refer to [this document](https://help.github.com/articles/creating-a-pull-request/) for detailed steps on how to create a pull request. After a Pull Request gets peer reviewed and approved, it will be merged.

## FAQs
### How does puml-c4-to-object-go manages dependencies?
``puml-c4-to-object-go`` uses `go mod` to manage its dependencies.
- Run `go get foo/bar` in the source folder to add the dependency to `go.mod` file.

To remove a dependency
- Edit your code and remove the import reference.
- Run `go mod tidy` in the source folder to remove dependency from `go.mod` file.

### What are the coding guidelines?
Refer: [Effective Go](https://github.com/golang/go/wiki/CodeReviewComments) article from Golang project. If you observe offending code, please feel free to send a pull request.