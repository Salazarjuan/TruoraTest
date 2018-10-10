'use strict';

var dbm;
var type;
var seed;

/**
  * We receive the dbmigrate dependency from dbmigrate initially.
  * This enables us to not have to rely on NODE_PATH.
  */
exports.setup = function(options, seedLink) {
  dbm = options.dbmigrate;
  type = dbm.dataType;
  seed = seedLink;
};

exports.up = function(db) {
  return db.createTable('recipesV1',{
      id:{type:'int', primaryKey:true, autoIncrement:true},
      name:{type:'string', notNull:true},
      description:{type:'string', notNull:false},
      oven:{type:'boolean', notNull:false},
      time:{type:'int', notNull:false},
      noPersons:{type:'int', notNull:false}
    });
};

exports.down = function(db) {
return db.dropTable('recipesV1');
};


exports._meta = {
  "version": 1
};
