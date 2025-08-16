(* let rec hot_maker(x) = *)
(*   function *)
(*     (x) *)
(*     -> Hot(x) *)

type seasoning = Salt | Pepper

type num =
  | Zero
  | One_more_than of num

type 'a open_faced_sandwich =
  | Bread of 'a
  | Slice of 'a open_faced_sandwich

type shish_kebab =
  | Skewer
  | Onion of shish_kebab
  | Lamb of shish_kebab
  | Tomato of shish_kebab

let rec only_onions = function
  | Skewer -> true
  | Onion(x) -> only_onions x
  | _ -> false

(* type assertion *)
let _ = (only_onions: shish_kebab -> bool)

let rec is_vegetarian = function
  | Skewer -> true
  | Lamb(_) -> false
  | Onion(x) | Tomato(x) -> is_vegetarian x


