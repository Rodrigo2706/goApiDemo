
DROP DATABASE IF EXISTS GoApiTutorial;

CREATE DATABASE GoApiTutorial;

USE GoApiTutorial;

CREATE TABLE TUT_Users (
    UserId SERIAL PRIMARY KEY,
    Name NVARCHAR(80) NOT NULL,
    Lastname NVARCHAR(80) NOT NULL,
    Email NVARCHAR(100) NOT NULL,
    Enabled BIT DEFAULT 1,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP PROCEDURE IF EXISTS TUTSP_GetUsers;

DELIMITER //

CREATE PROCEDURE TUTSP_GetUsers(pOffset INT, pLimit INT)
BEGIN

	DECLARE vOffsetValue INT;
    
    SET vOffsetValue = pLimit * pOffset;

	SELECT Name, Lastname, Email
	FROM TUT_Users
    WHERE Enabled = 1
	LIMIT pLimit 
    OFFSET vOffsetValue;

END //

DELIMITER ;

DROP PROCEDURE IF EXISTS TUTSP_CreateUser;

DELIMITER //

CREATE PROCEDURE TUTSP_CreateUser(pName NVARCHAR(80), pLastname NVARCHAR(80), pEmail NVARCHAR(100))
BEGIN

	IF EXISTS(SELECT UserId FROM TUT_Users WHERE Email = pEmail AND Enabled = 1) THEN
		SELECT 10000 returnCode; -- Email already exists
	ELSE
    BEGIN

		INSERT INTO TUT_Users(Name, Lastname, Email)
		VALUES(pName, pLastName, pEmail);
		
		SELECT 200 returnCode;
		
    END;
    END IF;
END //

DELIMITER ;

DROP PROCEDURE IF EXISTS TUTSP_GetUserByEmail;

DELIMITER //

CREATE PROCEDURE TUTSP_GetUserByEmail(pEmail NVARCHAR(100))
BEGIN

	IF NOT EXISTS(SELECT UserId FROM TUT_Users WHERE Email = pEmail AND Enabled = 1) THEN
		SELECT 10001 returnCode; -- User not found
    ELSE
    BEGIN
		SELECT Name, Lastname, Email
		FROM TUT_Users
		WHERE Email = pEmail AND Enabled = 1;
	END;
    END IF;

END //

DELIMITER ;

DROP PROCEDURE IF EXISTS TUTSP_UpdateUserByEmail;

DELIMITER //

CREATE PROCEDURE TUTSP_UpdateUserByEmail(pEmail NVARCHAR(100), pName NVARCHAR(80), pLastname NVARCHAR(80))
BEGIN

	IF NOT EXISTS(SELECT UserId FROM TUT_Users WHERE Email = pEmail AND Enabled = 1) THEN
		SELECT 10001 returnCode; -- User not found
    ELSE
    BEGIN

		UPDATE TUT_Users
        SET Name = COALESCE(pName, Name),
			Lastname = COALESCE(pLastname, Lastname)
		WHERE Email = pEmail AND Enabled = 1;
		
		SELECT 200 returnCode;
		
    END;
    END IF;
END //

DELIMITER ;

DROP PROCEDURE IF EXISTS TUTSP_DeleteUserByEmail;

DELIMITER //

CREATE PROCEDURE TUTSP_DeleteUserByEmail(pEmail NVARCHAR(100))
BEGIN

	IF NOT EXISTS(SELECT UserId FROM TUT_Users WHERE Email = pEmail AND Enabled = 1) THEN
		SELECT 10001 returnCode; -- User not found
    ELSE
    BEGIN

		UPDATE TUT_Users
        SET Enabled = 0
		WHERE Email = pEmail AND Enabled = 1;
		
		SELECT 200 returnCode;
		
    END;
    END IF;
END //

DELIMITER ;



