#!/usr/bin/env ruby

require 'thor'
require 'thor/rake_compat'
require 'rspec/core/rake_task'

module Go
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
  class Rake < Thor
    include Thor::RakeCompat

    desc 'rspec', 'Run RSpec tests. like there are any'
    def rspec
      RSpec::Core::RakeTask.new(:spec) do |t|
        t.spec_opts = ['--color']
        t.spec_files = FileList['spec/**/*_spec.rb']
      end
    end
  end
end
