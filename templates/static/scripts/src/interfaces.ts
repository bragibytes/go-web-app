
export interface user {
    _id?:string
    name?:string
    email?:string
    password?:string
    confirm_password?:string
    bio?:string
    created_at?:Date
    updated_at?:Date
}
export interface server_response {
    message_type:string
    message:string
    data?:any
}
export interface post {
    _id?:string
    _author?:string
    title?:string
    content?:string
    created_at?:Date
    updated_at?:Date
    score?:number
}

export interface vote {
    _id?:string
    _author?:string
    _parent?:string
    value?:number
    created_at?:Date
    updated_at?:Date
}
export interface comment {
    _id?:string
    _author?:string
    _parent?:string
    content?:string
    created_at?:Date
    updated_at?:Date
    score?:number
}
// type Comment struct {
// 	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
// 	Author    primitive.ObjectID `json:"_author" bson:"_author"`
// 	Parent    primitive.ObjectID `json:"_parent" bson:"_parent"`
// 	Content   string             `json:"content" bson:"content" validate:"required" min:"3" max:"255"`
// 	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
// 	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
// 	Score     int32              `json:"score" bson:"score,omitempty"`
// }
// type Vote struct {
// 	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Author    primitive.ObjectID `json:"_author" bson:"_author"`
// 	Parent    primitive.ObjectID `json:"_parent" bson:"_parent"`
// 	IsUpvote  bool               `json:"is_upvote" bson:"is_upvote"`
// 	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
// 	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
// }

// type Post struct {
//     ID         primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
//     Author     primitive.ObjectID   `json:"_author" bson:"_author,omitempty"`
//     Title      string               `json:"title" bson:"title,omitempty" validate:"required" min:"3" max:"200"`
//     Content    string               `json:"content" bson:"content,omitempty" validate:"required" min:"5" max:"10000"`
//     CommentIDs []primitive.ObjectID `json:"comments" bson:"comments,omitempty"`
//     CreatedAt  time.Time            `json:"created_at" bson:"created_at,omitempty"`
//     UpdatedAt  time.Time            `json:"updated_at" bson:"updated_at,omitempty"`
//
//     Score int32 `json:"score" bson:"-"`
// }

// ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
// Name            string             `json:"name" bson:"name,omitempty" validate:"required, gt=3"`
// Email           string             `json:"email" bson:"email,omitempty" validate:"required,email"`
// Password        string             `json:"password,omitempty" bson:"password,omitempty" validate:"required, gt=8"`
// ConfirmPassword string             `json:"confirmPassword,omitempty" bson:"-" validate:"required,eqfield=Password"`
// Bio             string             `json:"bio,omitempty" bson:"bio,omitempty"`
// CreatedAt       time.Time          `json:"created_at" bson:"created_at,omitempty"`
// UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
//
// Stats stats `json:"stats,omitempty" bson:"stats,omitempty"`