CREATE TABLE public.author(
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book(
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name VARCHAR(100) NOT NULL,
                            author_id UUID NOT NULL,
                            CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('People'); --bd5f45e1-cfbc-434f-aeb2-1148f84a5af0
INSERT INTO author (name) VALUES ('Jack London'); --bed798ce-78cc-4f96-90f8-fe20639dc159
INSERT INTO author (name) VALUES ('Shipwreck'); --add94412-88af-407c-8843-d62eea79739c

INSERT INTO book (name, author_id) VALUES ('Kolobok','bd5f45e1-cfbc-434f-aeb2-1148f84a5af0');
INSERT INTO book (name, author_id) VALUES ('Martin Iden','bed798ce-78cc-4f96-90f8-fe20639dc159');
INSERT INTO book (name, author_id) VALUES ('Romeo and Juliet','add94412-88af-407c-8843-d62eea79739c');