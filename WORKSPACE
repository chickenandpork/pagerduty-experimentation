load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

#
# GO and Rust are the new hotness; let's build some Go
#
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2d536797707dd1697441876b2e862c58839f975c8fc2f0f96636cbd428f45866",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.5/rules_go-v0.23.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.5/rules_go-v0.23.5.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

# Gazelle.  Like when people needed help to write USL makefiles, but it's also an opinionated formatter of BUILD files.
http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

http_archive(
    name = "com_github_atlassian_bazel_tools",
    sha256 = "a1ebb48984f0a8c38c5ab155043f9ed63a3880ef53df9785154ebab946f4821f",
    strip_prefix = "bazel-tools-541e546def26783063c7630495a618d48be2b4cb",
    urls = [
        "https://github.com/atlassian/bazel-tools/archive/541e546def26783063c7630495a618d48be2b4cb.tar.gz",
    ],
)

load(":bzl/go_repositories.bzl", "go_repositories")

# gazelle:repository_macro bzl/go_repositories.bzl%go_repositories
go_repositories()


# vim: ts=8 et sw=4 sts=4 :
