-- CreateTable
CREATE TABLE "games" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "player_id" UUID,
    "board" CHAR(1)[],
    "begin_date" DATE,

    CONSTRAINT "games_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "username" VARCHAR(20) NOT NULL,
    "password" TEXT NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "users_username_key" ON "users"("username");

-- AddForeignKey
ALTER TABLE "games" ADD CONSTRAINT "games_player_id_fkey" FOREIGN KEY ("player_id") REFERENCES "users"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

