CREATE TABLE `organisation` (
	`org_id` integer PRIMARY KEY NOT NULL,
	`name` text
);
--> statement-breakpoint
CREATE TABLE `team` (
	`team_id` integer PRIMARY KEY NOT NULL,
	`name` text,
	`org_id` integer
);
--> statement-breakpoint
CREATE TABLE `user` (
	`user_id` integer PRIMARY KEY NOT NULL,
	`name` text,
	`team_id` integer
);
