#!/usr/bin/env ruby
#
#
#

task default: ['go:fmt']

namespace :go do
  desc 'Format all Golang files.'
  task :fmt do
    Dir.glob('./**/*.go') do |go_file|
      next if File.directory?(go_file)
      puts '=> Running go fmt'
      sh "go fmt #{go_file}"
    end
  end

  desc 'Build the current Golang project.'
  task build: %w(go:fmt) do
    sh 'go build'
  end

  desc 'Run the current main.go file.'
  task run: %w(go:fmt) do
    sh 'go run main.go'
  end

  desc 'Run all the tests in current directory.'
  task test: %w(go:fmt) do
    sh 'go test'
  end

  desc 'Install the project to $GOBIN'
  task install: %w(go:fmt) do
    sh 'go install'
  end
end
