package customresource

const actionsetCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: actionsets.cr.kanister.io
spec:
  group: cr.kanister.io
  names:
    kind: ActionSet
    listKind: ActionSetList
    plural: actionsets
    singular: actionset
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              actions:
                items:
                  description: ActionSpec is the specification for a single Action.
                  properties:
                    artifacts:
                      additionalProperties:
                        properties:
                          keyValue:
                            additionalProperties:
                              type: string
                            type: object
                          kopiaSnapshot:
                              type: string
                        type: object
                      description: Artifacts will be passed as inputs into this phase.
                      type: object
                    blueprint:
                      description: Blueprint with instructions on how to execute this
                        action.
                      type: string
                    configMaps:
                      additionalProperties:
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          group:
                            description: API Group of the referent.
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                            type: string
                          resource:
                            description: Resource name of the referent.
                            type: string
                        type: object
                      description: ConfigMaps that we'll get and pass into the blueprint.
                      type: object
                    name:
                      description: 'Name is the action we''ll perform. For example:
                        backup or restore.'
                      type: string
                    object:
                      description: Object refers to the thing we'll perform this action
                        on.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        group:
                          description: API Group of the referent.
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                          type: string
                        resource:
                          description: Resource name of the referent.
                          type: string
                      type: object
                    options:
                      additionalProperties:
                        type: string
                      description: Options will be used to specify additional values
                        to be used in the Blueprint.
                      type: object
                    podOverride:
                      type: object
                      description: PodOverride is used to specify pod specs that will
                        override the default pod specs
                    preferredVersion:
                      description: PreferredVersion will be used to select the preferred
                        version of Kanister functions to be executed for this action
                      type: string
                    profile:
                      description: Profile is use to specify the location where store
                        artifacts and the credentials authorized to access them.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        group:
                          description: API Group of the referent.
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                          type: string
                        resource:
                          description: Resource name of the referent.
                          type: string
                      type: object
                    secrets:
                      additionalProperties:
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          group:
                            description: API Group of the referent.
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                            type: string
                          resource:
                            description: Resource name of the referent.
                            type: string
                        type: object
                      description: Secrets that we'll get and pass into the blueprint.
                      type: object
                  type: object
                type: array
            type: object
          status:
            description: ActionSetStatus is the status for the actionset. This should
              only be updated by the controller.
            properties:
              actions:
                items:
                  properties:
                    artifacts:
                      additionalProperties:
                        properties:
                          keyValue:
                            additionalProperties:
                              type: string
                            type: object
                        type: object
                      description: Artifacts created by this phase.
                      type: object
                    blueprint:
                      description: Blueprint with instructions on how to execute this
                        action.
                      type: string
                    name:
                      description: 'Name is the action we''ll perform. For example:
                        backup or restore.'
                      type: string
                    object:
                      description: Object refers to the thing we'll perform this action
                        on.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        group:
                          description: API Group of the referent.
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: http://kubernetes.io/docs/user-guide/namespaces'
                          type: string
                        resource:
                          description: Resource name of the referent.
                          type: string
                      type: object
                    phases:
                      description: Phases are sub-actions an are executed sequentially.
                      items:
                        properties:
                          name:
                            type: string
                          output:
                            x-kubernetes-preserve-unknown-fields: true
                            type: object
                          state:
                            type: string
                        type: object
                      type: array
                  type: object
                type: array
              error:
                properties:
                  message:
                    type: string
                type: object
              state:
                type: string
            type: object
        type: object
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
