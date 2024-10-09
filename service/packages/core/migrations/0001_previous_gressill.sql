PRAGMA foreign_keys=OFF;--> statement-breakpoint
CREATE TABLE `__new_organisation` (
	`org_id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text
);
--> statement-breakpoint
INSERT INTO `__new_organisation`("org_id", "name") SELECT "org_id", "name" FROM `organisation`;--> statement-breakpoint
DROP TABLE `organisation`;--> statement-breakpoint
ALTER TABLE `__new_organisation` RENAME TO `organisation`;--> statement-breakpoint
PRAGMA foreign_keys=ON;--> statement-breakpoint
CREATE TABLE `__new_team` (
	`team_id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text,
	`org_id` integer
);
--> statement-breakpoint
INSERT INTO `__new_team`("team_id", "name", "org_id") SELECT "team_id", "name", "org_id" FROM `team`;--> statement-breakpoint
DROP TABLE `team`;--> statement-breakpoint
ALTER TABLE `__new_team` RENAME TO `team`;--> statement-breakpoint
CREATE TABLE `__new_user` (
	`user_id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text,
	`team_id` integer
);
--> statement-breakpoint
INSERT INTO `__new_user`("user_id", "name", "team_id") SELECT "user_id", "name", "team_id" FROM `user`;--> statement-breakpoint
DROP TABLE `user`;--> statement-breakpoint
ALTER TABLE `__new_user` RENAME TO `user`;