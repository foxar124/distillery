# Aliases

You can configure aliases for your installation sources. This is useful if you don't want to type the whole
path all the time.

## Simple Definition

=== "YAML"

    ```yaml
    aliases:
      dist: ekristen/distillery
      aws-nuke: ekristen/aws-nuke
      age: filosottile/age
    ```

=== "TOML"

    ```toml
    [aliases]
    dist = "ekristen/distillery"
    aws-nuke = "ekristen/aws-nuke"
    age = "filosottile/age"
    ```

## With Version

=== "YAML"

    ```yaml
    aliases:
      dist: ekristen/distillery
      aws-nuke: ekristen/aws-nuke
      age: filosottile/age@1.0.0
    ```

=== "TOML"

    ```toml
    [aliases]
    dist = "ekristen/distillery"
    aws-nuke = "ekristen/aws-nuke"
    age = "filosottile/age@1.0.0"
    ```

## With Version as Object

=== "YAML"

    ```yaml
    aliases:
      dist: ekristen/distillery
      aws-nuke: ekristen/aws-nuke
      age:
        name: filosottile/age
        version: 1.0.0
    ```

=== "TOML"

    ```toml
    [aliases]
    dist = "ekristen/distillery"
    aws-nuke = "ekristen/aws-nuke"
    
    [aliases.age]
    name = "filosottile/age"
    version = "1.0.0"
    ```

## With Providers

=== "YAML"


    ```yaml
    aliases:
      age: github/filosottile/age
      gitlab-runner: gitlab/gitlab-org/gitlab-runner
    ```

=== "TOML"

    ```toml
    [aliases]
    age = "github/filosottile/age"
    gitlab-runner = "gitlab/gitlab-org/gitlab-runner"
    ```