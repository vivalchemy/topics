TRUNCATE organizations CASCADE;
INSERT INTO organizations (id, name) VALUES
('ce17ddfa-1a31-47c5-abb9-c43c8dfd5df5', 'OpenAI'),
('4763aacd-faf5-4ec3-bbc0-a6b7172ce440', 'Tech Innovators'),
('62b5365b-a0c3-466e-8c40-9b6b7d016213', 'Data Wizards'),
('e2ae29b5-8fcb-4cea-a780-1f3fcf88008f', 'AI Pioneers'),
('bc5b9c77-7fec-4d83-a66b-199aa2598230', 'Quantum Computing');

TRUNCATE users CASCADE;
INSERT INTO users (id, name, username, password) VALUES
('aa2412c8-ca01-431b-a2a8-b1d240d09d46', 'Alice', 'alice1', 'password1'),
('50acf1af-3801-440c-8a34-0206937c0baa', 'Bob', 'bob2', 'password2'),
('c69b05da-b19e-4dcc-9d7c-feea6b2d4eda', 'Charlie', 'charlie3', 'password3'),
('922bbde5-9355-460c-a681-9c7b8b97b1ea', 'David', 'david4', 'password4'),
('a7ace2e7-e6fa-4b85-b46f-ebb352f711c1', 'Eve', 'eve5', 'password5'),
('4f3217a3-da8a-49d2-8a21-392d304edae9', 'Frank', 'frank6', 'password6'),
('616f17ac-d630-4a35-bebb-695a729f28fd', 'Grace', 'grace7', 'password7'),
('7eb02f4a-8322-4e45-84ac-56f42e1ee5b4', 'Hank', 'hank8', 'password8'),
('2688195b-51ef-41f9-a77f-181fa8476ac3', 'Ivy', 'ivy9', 'password9'),
('ec8760a6-17d8-4bf3-bfe9-3c9887f7ffb2', 'Jack', 'jack10', 'password10'),
('31efbe31-9654-40a8-bde6-b08d603940f8', 'Ken', 'ken11', 'password11'),
('235f0762-10b6-4443-9480-06dc7fdd9f96', 'Laura', 'laura12', 'password12'),
('eeffd535-5c17-43b9-ae84-2da206628c7e', 'Mike', 'mike13', 'password13'),
('f7e06048-3400-4d76-ad8e-443710b104cb', 'Nina', 'nina14', 'password14'),
('20d69efb-7df6-46d0-80bf-cd5ab51d79f2', 'Oscar', 'oscar15', 'password15'),
('5d597027-ca21-4b41-bef2-92462e57cc55', 'Patty', 'patty16', 'password16'),
('01115fc9-d789-4729-a842-9a2c01d4b324', 'Quinn', 'quinn17', 'password17'),
('6ee4d2af-da34-4b23-acf7-029c46e87383', 'Rick', 'rick18', 'password18'),
('092bcea5-ea6c-47c8-9916-c43e96be2c03', 'Sara', 'sara19', 'password19'),
('a0cd345b-35d1-4dfd-a009-f2b582c55048', 'Tom', 'tom20', 'password20'),
('47fd3a70-edca-4c43-9de7-7f92199c1345', 'Uma', 'uma21', 'password21'),
('107cc425-83f5-4e70-b9d6-3c92781abd2e', 'Victor', 'victor22', 'password22'),
('4993a585-586e-40de-809a-5e0057f2214c', 'Wendy', 'wendy23', 'password23'),
('0ed74d20-3891-416b-aa20-c450ab06e5a7', 'Xander', 'xander24', 'password24'),
('9f390575-7ba2-4c11-8b61-760b30fe01a0', 'Yara', 'yara25', 'password25'),
('15fb4251-a9e1-4337-8baa-70710e5b2761', 'Zane', 'zane26', 'password26'),
('ea9b0955-7963-450d-91eb-b01df4343266', 'Alice', 'alice27', 'password27'),
('efc3eccf-a645-4f99-8dba-df84345d01b9', 'Bob', 'bob28', 'password28'),
('952002cd-54af-40ec-bc3f-0ff254e50bdc', 'Charlie', 'charlie29', 'password29'),
('1050a821-2515-49ef-b60a-facfbe852bac', 'David', 'david30', 'password30');

TRUNCATE organization_user CASCADE;
INSERT INTO organization_user (organization_id, user_id) VALUES
('ce17ddfa-1a31-47c5-abb9-c43c8dfd5df5', 'aa2412c8-ca01-431b-a2a8-b1d240d09d46'),
('ce17ddfa-1a31-47c5-abb9-c43c8dfd5df5', '107cc425-83f5-4e70-b9d6-3c92781abd2e'),
('ce17ddfa-1a31-47c5-abb9-c43c8dfd5df5', 'efc3eccf-a645-4f99-8dba-df84345d01b9'),
('bc5b9c77-7fec-4d83-a66b-199aa2598230', 'a0cd345b-35d1-4dfd-a009-f2b582c55048'),
('bc5b9c77-7fec-4d83-a66b-199aa2598230', 'aa2412c8-ca01-431b-a2a8-b1d240d09d46'),
('e2ae29b5-8fcb-4cea-a780-1f3fcf88008f', 'a7ace2e7-e6fa-4b85-b46f-ebb352f711c1'),
('e2ae29b5-8fcb-4cea-a780-1f3fcf88008f', 'aa2412c8-ca01-431b-a2a8-b1d240d09d46'),
('bc5b9c77-7fec-4d83-a66b-199aa2598230', '1050a821-2515-49ef-b60a-facfbe852bac'),
('62b5365b-a0c3-466e-8c40-9b6b7d016213', '107cc425-83f5-4e70-b9d6-3c92781abd2e'),
('4763aacd-faf5-4ec3-bbc0-a6b7172ce440', 'f7e06048-3400-4d76-ad8e-443710b104cb');
