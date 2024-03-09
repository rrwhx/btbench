/*
  Copyright (C) 2014, Daishi Kato <daishi@axlight.com>
  Copyright (C) 2014, Etienne Rossignon
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are

    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in th
      documentation and/or other materials provided with the distributio

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
  "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
  LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
  A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
  HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
  SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
  LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
  DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
  THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/* global BenchmarkSuite: false */

var vm = require('vm');
var fs = require('fs');
var path = require('path');
var os = require('os');

global.print = function(str) {
  console.log(str);
};

global.read = function(a, b) {
  var a = path.normalize(a);
  var c = fs.readFileSync(a);
  if (!c && a != path.resolve(a)) {
    a = path.join(__dirname, '..', 'src', a);
    c = fs.readFileSync(a);
  }
  if (c && !b) {
    c = c.toString();
  }
  return c;
};

function load(filename) {
  vm.runInThisContext(fs.readFileSync(filename, 'utf8'), filename);
}


var benchs = {
  "Richards" : 120000,
  "DeltaBlue" : 150000,
  "Crypto" : 5000,
  "RayTrace" : 20000,
  "EarleyBoyer" : 3000,
  "RegExp" : 1000,
  "Splay" : 50000,
  "NavierStokes" : 3000,
  "PdfJS" : 500,
  "Mandreel" : 400,
  "Gameboy" : 100,
  "CodeLoad" : 2000,
  "Box2D" : 3000,
  "zlib" : 60,
  "Typescript" : 60,
}

if (process.argv[2]) {
  // print(process.argv)
  var text = process.argv[2]
  try {
    benchs = JSON.parse(text)
  } catch (e) {
    if (benchs[text]) {
      var bench = {}
      bench[text] = benchs[text]
      benchs = bench
    } else {
      print("Usage : node octane.js '{\"Typescript\":10}'")
      print("Usage : node octane.js Typescript")
      process.exit(0)
    }
  }
}

var base_dir = __dirname + '/octane/';
load(__dirname + '/base.js');
if (benchs["Richards"])     load(base_dir + 'richards.js');
if (benchs["DeltaBlue"])    load(base_dir + 'deltablue.js');
if (benchs["Crypto"])       load(base_dir + 'crypto.js');
if (benchs["RayTrace"])     load(base_dir + 'raytrace.js');
if (benchs["EarleyBoyer"])  load(base_dir + 'earley-boyer.js');
if (benchs["RegExp"])       load(base_dir + 'regexp.js');
if (benchs["Splay"])        load(base_dir + 'splay.js');
if (benchs["NavierStokes"]) load(base_dir + 'navier-stokes.js');
if (benchs["PdfJS"])        load(base_dir + 'pdfjs.js');
if (benchs["Mandreel"])     load(base_dir + 'mandreel.js');
if (benchs["Gameboy"])      load(base_dir + 'gbemu-part1.js');
if (benchs["Gameboy"])      load(base_dir + 'gbemu-part2.js');
if (benchs["CodeLoad"])     load(base_dir + 'code-load.js');
if (benchs["Box2D"])        load(base_dir + 'box2d.js');
if (benchs["zlib"])         load(base_dir + 'zlib.js');
if (benchs["zlib"])         load(base_dir + 'zlib-data.js');
if (benchs["Typescript"])   load(base_dir + 'typescript.js');
if (benchs["Typescript"])   load(base_dir + 'typescript-input.js');
if (benchs["Typescript"])   load(base_dir + 'typescript-compiler.js');

// BenchmarkSuite.config.doWarmup = undefined;
// BenchmarkSuite.config.doDeterministic = undefined;

console.log('    hostname     :', os.hostname());
console.log('    node version :', process.version);
console.log('      V8 version :', process.versions['v8']);
console.log(' platform & arch :', process.platform, process.arch);
console.log('');
console.log(' config :', require('util').inspect(process.config, {
  colors: true,
  depth: 10
}));
console.log('');

BenchmarkSuite.RunSuites(benchs);
console.log(' duration ', process.uptime(), ' seconds');
