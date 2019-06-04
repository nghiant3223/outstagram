# Create Story

```plantuml
@startuml
skinparam sequenceArrowThickness 2
skinparam roundcorner 5
skinparam ParticipantPadding 50

actor User
participant "PostController" as A
participant "ImageService" as B
participant "ImageUtil" as C
participant "PostImageService" as E
participant "PostService" as D

User -> A: createPostHandler(contet, images)
activate A

loop for every image

A -> B: newImage(image)
activate B

B->C: createThumbnails(image)
activate C

C-->B: thumbnailURLs: []string
deactivate C

B->B: saveImage(thumbnailURLs)

B-->A: imageID
deactivate B

A->E: newPostImage(imageID)
activate E

E->E: setImageID(imageID)

E->E: setReactableID()

E->E: setCommentableID()

E->E: setViewableID()

E-->A: PostImage{}

end

A->D: savePost(content, postImages)
activate D

D-->A: Post{}

@enduml
```
