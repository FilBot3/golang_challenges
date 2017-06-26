#!/usr/bin/env ruby

require 'thor'

module Golang
  class Actions < Thor
    desc 'fmt', 'Format all Golang files in project.'
    def fmt
      Dir.glob('./**/*.go') do |go_file|
        next if File.directory?(go_file)
        puts '=> Running go fmt'
        sh "go fmt #{go_file}"
      end
    end
  end
end
