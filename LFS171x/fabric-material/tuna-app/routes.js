//SPDX-License-Identifier: Apache-2.0

var exames = require('./controller.js');

module.exports = function(app){

  app.get('/get_tuna/:id', function(req, res){
    exames.get_tuna(req, res);
  });
  app.get('/add_tuna/:tuna', function(req, res){
    exames.add_tuna(req, res);
  });
  app.get('/get_all_exames', function(req, res){
    exames.get_all_exames(req, res);
  });
  app.get('/change_holder/:holder', function(req, res){
    exames.change_holder(req, res);
  });
}
