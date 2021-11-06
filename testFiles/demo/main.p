RUN firstProc.
RUN secondProc.

PROCEDURE firstProc:
    RUN secondProc.
END PROCEDURE.

PROCEDURE secondProc:
    RUN thridProc.
END PROCEDURE.

PROCEDURE thridProc:
    MESSAGE "This is a third procedure".
END PROCEDURE.