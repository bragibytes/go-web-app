
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
    message_type:number
    message:string
    data?:any
    code:number
}

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