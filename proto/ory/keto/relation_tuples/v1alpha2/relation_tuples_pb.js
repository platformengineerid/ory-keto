// @generated by protoc-gen-es v1.7.2 with parameter "target=js,js_import_style=module"
// @generated from file ory/keto/relation_tuples/v1alpha2/relation_tuples.proto (package ory.keto.relation_tuples.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * RelationTuple defines a relation between an Object and a Subject.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.RelationTuple
 */
export const RelationTuple = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.RelationTuple",
  () => [
    { no: 1, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "object", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "relation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "subject", kind: "message", T: Subject },
  ],
);

/**
 * The query for listing relationships.
 * Clients can specify any optional field to
 * partially filter for specific relationships.
 *
 * Example use cases (namespace is always required):
 *  - object only: display a list of all permissions referring to a specific object
 *  - relation only: get all groups that have members; get all directories that have content
 *  - object & relation: display all subjects that have a specific permission relation
 *  - subject & relation: display all groups a subject belongs to; display all objects a subject has access to
 *  - object & relation & subject: check whether the relation tuple already exists
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.RelationQuery
 */
export const RelationQuery = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.RelationQuery",
  () => [
    { no: 1, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "object", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "relation", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "subject", kind: "message", T: Subject, opt: true },
  ],
);

/**
 * Subject is either a concrete subject id or
 * a `SubjectSet` expanding to more Subjects.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.Subject
 */
export const Subject = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.Subject",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "ref" },
    { no: 2, name: "set", kind: "message", T: SubjectSet, oneof: "ref" },
  ],
);

/**
 * SubjectSet refers to all subjects who have
 * the same `relation` on an `object`.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.SubjectSet
 */
export const SubjectSet = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.SubjectSet",
  () => [
    { no: 1, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "object", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "relation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

