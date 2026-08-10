using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace dotnetservice.DataAccess.Migrations
{
    /// <inheritdoc />
    public partial class UpdateUserTableName : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_file_counter_user_user_id",
                table: "file_counter");

            migrationBuilder.DropPrimaryKey(
                name: "PK_user",
                table: "user");

            migrationBuilder.RenameTable(
                name: "user",
                newName: "app_user");

            migrationBuilder.RenameIndex(
                name: "IX_user_password_email",
                table: "app_user",
                newName: "IX_app_user_password_email");

            migrationBuilder.RenameIndex(
                name: "IX_user_email",
                table: "app_user",
                newName: "IX_app_user_email");

            migrationBuilder.AddPrimaryKey(
                name: "PK_app_user",
                table: "app_user",
                column: "id");

            migrationBuilder.AddForeignKey(
                name: "FK_file_counter_app_user_user_id",
                table: "file_counter",
                column: "user_id",
                principalTable: "app_user",
                principalColumn: "id",
                onDelete: ReferentialAction.Cascade);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_file_counter_app_user_user_id",
                table: "file_counter");

            migrationBuilder.DropPrimaryKey(
                name: "PK_app_user",
                table: "app_user");

            migrationBuilder.RenameTable(
                name: "app_user",
                newName: "user");

            migrationBuilder.RenameIndex(
                name: "IX_app_user_password_email",
                table: "user",
                newName: "IX_user_password_email");

            migrationBuilder.RenameIndex(
                name: "IX_app_user_email",
                table: "user",
                newName: "IX_user_email");

            migrationBuilder.AddPrimaryKey(
                name: "PK_user",
                table: "user",
                column: "id");

            migrationBuilder.AddForeignKey(
                name: "FK_file_counter_user_user_id",
                table: "file_counter",
                column: "user_id",
                principalTable: "user",
                principalColumn: "id",
                onDelete: ReferentialAction.Cascade);
        }
    }
}
